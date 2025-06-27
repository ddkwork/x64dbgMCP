cmake_minimum_required(VERSION 3.15)
project(MCPx64dbg)

# set c++ standard
set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

# define architecture properties
if(CMAKE_SIZEOF_VOID_P EQUAL 4)
    set(arch_bits 32)
    set(arch_suffix "_x86")
    set(plugin_ext ".dp32")
elseif(CMAKE_SIZEOF_VOID_P EQUAL 8)
    set(arch_bits 64)
    set(arch_suffix "_x64")
    set(plugin_ext ".dp64")
endif()

# define plugin sdk directory and project root
set(pluginsdk_dir "${CMAKE_SOURCE_DIR}/pluginsdk")
set(project_root "${CMAKE_SOURCE_DIR}")

# include directories
include_directories(
    "${project_root}"
    "${pluginsdk_dir}"
)

# add source files
file(GLOB sources "*.cpp")

# create shared library (plugin)
add_library(${PROJECT_NAME} SHARED ${sources})
set_target_properties(${PROJECT_NAME} PROPERTIES
    PREFIX ""
    OUTPUT_NAME "${PROJECT_NAME}"
    SUFFIX "${plugin_ext}"
)

# function to find libraries with consistent path structure
function(find_plugin_sdk_lib lib_name output_var)
    set(lib_path "${pluginsdk_dir}/${lib_name}/${lib_name}${arch_suffix}.lib")
    set(fallback_path "${pluginsdk_dir}/${lib_name}/${lib_name}.lib")
    
    if(EXISTS "${lib_path}")
        set(${output_var} "${lib_path}" PARENT_SCOPE)
    elseif(EXISTS "${fallback_path}")
        message(STATUS "using fallback: ${fallback_path}")
        set(${output_var} "${fallback_path}" PARENT_SCOPE)
    else()
        message(FATAL_ERROR "library ${lib_name} not found at ${lib_path} or ${fallback_path}")
    endif()
endfunction()

# find all required sdk libraries
find_plugin_sdk_lib(TitanEngine titan_engine_lib)
find_plugin_sdk_lib(DeviceNameResolver device_resolver_lib)
find_plugin_sdk_lib(jansson jansson_lib)
find_plugin_sdk_lib(lz4 lz4_lib)
find_plugin_sdk_lib(XEDParse xedparse_lib)

# determine debugger core libraries
if(arch_bits EQUAL 32)
    set(dbg_lib "${pluginsdk_dir}/x32dbg.lib")
    set(bridge_lib "${pluginsdk_dir}/x32bridge.lib")
else()
    set(dbg_lib "${pluginsdk_dir}/x64dbg.lib")
    set(bridge_lib "${pluginsdk_dir}/x64bridge.lib")
endif()

# link all dependencies in a single command
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

# post-build information
set(plugin_path "${CMAKE_BINARY_DIR}/${CMAKE_CFG_INTDIR}/${PROJECT_NAME}${plugin_ext}")
set(install_path "$ENV{ProgramFiles}/x64dbg/release/x${arch_bits}/plugins/${PROJECT_NAME}${plugin_ext}")

add_custom_command(TARGET ${PROJECT_NAME} POST_BUILD
    COMMAND ${CMAKE_COMMAND} -E echo ""
    COMMAND ${CMAKE_COMMAND} -E echo "=== plugin build complete ==="
    COMMAND ${CMAKE_COMMAND} -E echo "plugin built at: ${plugin_path}"
    COMMAND ${CMAKE_COMMAND} -E echo "install manually to: ${install_path}"
    COMMAND ${CMAKE_COMMAND} -E echo "=============================="
)
