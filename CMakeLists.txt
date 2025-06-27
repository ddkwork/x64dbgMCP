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

# Set paths to x64dbg SDK
set(X64DBG_SDK_PATH "${CMAKE_SOURCE_DIR}")

# Include directories
include_directories(
    ${X64DBG_SDK_PATH}
    ${X64DBG_SDK_PATH}/pluginsdk
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

# Function to find SDK libraries
function(find_sdk_lib LIB_NAME OUTPUT_VAR)
    if(EXISTS "${X64DBG_SDK_PATH}/pluginsdk/${LIB_NAME}/${LIB_NAME}${ARCH_SUFFIX}.lib")
        set(${OUTPUT_VAR} "${X64DBG_SDK_PATH}/pluginsdk/${LIB_NAME}/${LIB_NAME}${ARCH_SUFFIX}.lib" PARENT_SCOPE)
    else()
        message(STATUS "${LIB_NAME}${ARCH_SUFFIX}.lib not found, trying ${LIB_NAME}.lib")
        set(${OUTPUT_VAR} "${X64DBG_SDK_PATH}/pluginsdk/${LIB_NAME}/${LIB_NAME}.lib" PARENT_SCOPE)
    endif()
endfunction()

# Find all required SDK libraries
find_sdk_lib(TitanEngine TITAN_ENGINE_LIB)
find_sdk_lib(DeviceNameResolver DEVICE_RESOLVER_LIB)
find_sdk_lib(jansson JANSSON_LIB)
find_sdk_lib(lz4 LZ4_LIB)
find_sdk_lib(XEDParse XEDPARSE_LIB)

# 确定调试器核心库
if(ARCH_BITS EQUAL 32)
    set(DBG_LIB "${X64DBG_SDK_PATH}/pluginsdk/x32dbg.lib")
    set(BRIDGE_LIB "${X64DBG_SDK_PATH}/pluginsdk/x32bridge.lib")
else()
    set(DBG_LIB "${X64DBG_SDK_PATH}/pluginsdk/x64dbg.lib")
    set(BRIDGE_LIB "${X64DBG_SDK_PATH}/pluginsdk/x64bridge.lib")
endif()

# 统一链接所有依赖项
link_directories(${X64DBG_SDK_PATH}/pluginsdk)
target_link_libraries(${PROJECT_NAME}
    ws2_32
    winhttp
    ${TITAN_ENGINE_LIB}
    ${DEVICE_RESOLVER_LIB}
    ${JANSSON_LIB}
    ${LZ4_LIB}
    ${XEDPARSE_LIB}
    ${DBG_LIB}        # 添加调试器核心库
    ${BRIDGE_LIB}     # 添加桥接库
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
