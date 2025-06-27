cmake_minimum_required(VERSION 3.15)
project(MCPx64dbg)

# Set C++ standard
set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

# Define architecture properties
if(CMAKE_SIZEOF_VOID_P EQUAL 4)
    set(arch_bits 32)
    set(arch_suffix "_x86")
    set(plugin_ext ".dp32")
elseif(CMAKE_SIZEOF_VOID_P EQUAL 8)
    set(arch_bits 64)
    set(arch_suffix "_x64")
    set(plugin_ext ".dp64")
endif()

# 清晰定义插件sdk目录和项目根目录
set(pluginsdk_dir "${CMAKE_SOURCE_DIR}/pluginsdk")
set(project_root "${CMAKE_SOURCE_DIR}")

# 包含目录 - 明确区分项目文件和sdk文件
include_directories(
    "${project_root}"        # 项目自己的头文件
    "${pluginsdk_dir}"       # sdk头文件
)

# Add source files
file(GLOB sources "*.cpp")

# Create shared library (plugin)
add_library(${PROJECT_NAME} SHARED ${sources})
set_target_properties(${PROJECT_NAME} PROPERTIES
    PREFIX ""
    OUTPUT_NAME "${PROJECT_NAME}"
    SUFFIX "${plugin_ext}"
)

# 统一库查找函数，使用清晰的路径结构
function(find_plugin_sdk_lib lib_name output_var)
    set(lib_path "${pluginsdk_dir}/${lib_name}/${lib_name}${arch_suffix}.lib")
    set(fallback_path "${pluginsdk_dir}/${lib_name}/${lib_name}.lib")
    
    if(EXISTS "${lib_path}")
        set(${output_var} "${lib_path}" PARENT_SCOPE)
    elseif(EXISTS "${fallback_path}")
        message(STATUS "Using fallback: ${fallback_path}")
        set(${output_var} "${fallback_path}" PARENT_SCOPE)
    else()
        message(FATAL_ERROR "Library ${lib_name} not found at ${lib_path} or ${fallback_path}")
    endif()
endfunction()

# Find all required sdk libraries
find_plugin_sdk_lib(TitanEngine titan_engine_lib)
find_plugin_sdk_lib(DeviceNameResolver device_resolver_lib)
find_plugin_sdk_lib(jansson jansson_lib)
find_plugin_sdk_lib(lz4 lz4_lib)
find_plugin_sdk_lib(XEDParse xedparse_lib)

# 确定调试器核心库
if(arch_bits EQUAL 32)
    set(dbg_lib "${pluginsdk_dir}/x32dbg.lib")
    set(bridge_lib "${pluginsdk_dir}/x32bridge.lib")
else()
    set(dbg_lib "${pluginsdk_dir}/x64dbg.lib")
    set(bridge_lib "${pluginsdk_dir}/x64bridge.lib")
endif()

# 统一链接所有依赖项
target_link_libraries(${PROJECT_NAME}
    ws2_32
    winhttp
    ${titan_engine_lib}
    ${device_resolver_lib}
    ${jansson_lib}
    ${lz4_lib}
    ${xedparse_lib}
    ${dbg_lib}
    ${bridge_lib}
)

# Post-build information
set(plugin_path "${CMAKE_BINARY_DIR}/${CMAKE_CFG_INTDIR}/${PROJECT_NAME}${plugin_ext}")
set(install_path "$ENV{ProgramFiles}/x64dbg/release/x${arch_bits}/plugins/${PROJECT_NAME}${plugin_ext}")

add_custom_command(TARGET ${PROJECT_NAME} POST_BUILD
    COMMAND ${CMAKE_COMMAND} -E echo ""
    COMMAND ${CMAKE_COMMAND} -E echo "=== Plugin build complete ==="
    COMMAND ${CMAKE_COMMAND} -E echo "Plugin built at: ${plugin_path}"
    COMMAND ${CMAKE_COMMAND} -E echo "Install manually to: ${install_path}"
    COMMAND ${CMAKE_COMMAND} -E echo "=============================="
)
