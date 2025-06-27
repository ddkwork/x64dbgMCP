cmake_minimum_required(VERSION 3.15)
project(MCPx64dbg)

# Set C++ standard
set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

# Define architecture properties
if(CMAKE_SIZEOF_VOID_P EQUAL 4)
    set(ARCH_BITS 32)
    set(ARCH_SUFFIX "_x86")
    set(PLUGIN_EXT ".dp32")
elseif(CMAKE_SIZEOF_VOID_P EQUAL 8)
    set(ARCH_BITS 64)
    set(ARCH_SUFFIX "_x64")
    set(PLUGIN_EXT ".dp64")
endif()

# 清晰定义插件SDK目录和项目根目录
set(PLUGINSDK_DIR "${CMAKE_SOURCE_DIR}/pluginsdk")
set(PROJECT_ROOT "${CMAKE_SOURCE_DIR}")

# 包含目录 - 明确区分项目文件和SDK文件
include_directories(
    "${PROJECT_ROOT}"        # 项目自己的头文件
    "${PLUGINSDK_DIR}"       # SDK头文件
)

# Add source files
file(GLOB SOURCES "*.cpp")

# Create shared library (plugin)
add_library(${PROJECT_NAME} SHARED ${SOURCES})
set_target_properties(${PROJECT_NAME} PROPERTIES
    PREFIX ""
    OUTPUT_NAME "${PROJECT_NAME}"
    SUFFIX "${PLUGIN_EXT}"
)

# 统一库查找函数，使用清晰的路径结构
function(find_plugin_sdk_lib LIB_NAME OUTPUT_VAR)
    set(LIB_PATH "${PLUGINSDK_DIR}/${LIB_NAME}/${LIB_NAME}${ARCH_SUFFIX}.lib")
    set(FALLBACK_PATH "${PLUGINSDK_DIR}/${LIB_NAME}/${LIB_NAME}.lib")
    
    if(EXISTS "${LIB_PATH}")
        set(${OUTPUT_VAR} "${LIB_PATH}" PARENT_SCOPE)
    elseif(EXISTS "${FALLBACK_PATH}")
        message(STATUS "Using fallback: ${FALLBACK_PATH}")
        set(${OUTPUT_VAR} "${FALLBACK_PATH}" PARENT_SCOPE)
    else()
        message(FATAL_ERROR "Library ${LIB_NAME} not found at ${LIB_PATH} or ${FALLBACK_PATH}")
    endif()
endfunction()

# Find all required SDK libraries
find_plugin_sdk_lib(TitanEngine TITAN_ENGINE_LIB)
find_plugin_sdk_lib(DeviceNameResolver DEVICE_RESOLVER_LIB)
find_plugin_sdk_lib(jansson JANSSON_LIB)
find_plugin_sdk_lib(lz4 LZ4_LIB)
find_plugin_sdk_lib(XEDParse XEDPARSE_LIB)

# 确定调试器核心库
if(ARCH_BITS EQUAL 32)
    set(DBG_LIB "${PLUGINSDK_DIR}/x32dbg.lib")
    set(BRIDGE_LIB "${PLUGINSDK_DIR}/x32bridge.lib")
else()
    set(DBG_LIB "${PLUGINSDK_DIR}/x64dbg.lib")
    set(BRIDGE_LIB "${PLUGINSDK_DIR}/x64bridge.lib")
endif()

# 统一链接所有依赖项
target_link_libraries(${PROJECT_NAME}
    ws2_32
    winhttp
    ${TITAN_ENGINE_LIB}
    ${DEVICE_RESOLVER_LIB}
    ${JANSSON_LIB}
    ${LZ4_LIB}
    ${XEDPARSE_LIB}
    ${DBG_LIB}
    ${BRIDGE_LIB}
)

# Post-build information
set(PLUGIN_PATH "${CMAKE_BINARY_DIR}/${CMAKE_CFG_INTDIR}/${PROJECT_NAME}${PLUGIN_EXT}")
set(INSTALL_PATH "$ENV{ProgramFiles}/x64dbg/release/x${ARCH_BITS}/plugins/${PROJECT_NAME}${PLUGIN_EXT}")

add_custom_command(TARGET ${PROJECT_NAME} POST_BUILD
    COMMAND ${CMAKE_COMMAND} -E echo ""
    COMMAND ${CMAKE_COMMAND} -E echo "=== Plugin build complete ==="
    COMMAND ${CMAKE_COMMAND} -E echo "Plugin built at: ${PLUGIN_PATH}"
    COMMAND ${CMAKE_COMMAND} -E echo "Install manually to: ${INSTALL_PATH}"
    COMMAND ${CMAKE_COMMAND} -E echo "=============================="
)
