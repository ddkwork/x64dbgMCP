cmake_minimum_required(VERSION 3.15)
project(MCPx64dbg)

# Set C++ standard
set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

# 定义架构属性
if(CMAKE_SIZEOF_VOID_P EQUAL 4)
    set(ARCH_BITS 32)
    set(ARCH_SUFFIX "_x86")
    set(PLUGIN_EXT ".dp32")
    set(ARCH_DEFINE "-D_WIN32")
elseif(CMAKE_SIZEOF_VOID_P EQUAL 8)
    set(ARCH_BITS 64)
    set(ARCH_SUFFIX "_x64")
    set(PLUGIN_EXT ".dp64")
    set(ARCH_DEFINE "-D_WIN64")
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

# 设置编译定义
target_compile_definitions(${PROJECT_NAME} PRIVATE ${ARCH_DEFINE})

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

# Link dependencies
link_directories(${X64DBG_SDK_PATH}/pluginsdk)
target_link_libraries(${PROJECT_NAME}
    ws2_32
    winhttp
    ${TITAN_ENGINE_LIB}
    ${DEVICE_RESOLVER_LIB}
    ${JANSSON_LIB}
    ${LZ4_LIB}
    ${XEDPARSE_LIB}
)

# Add architecture-specific dbg libraries
set(DBG_LIB_PATH "${X64DBG_SDK_PATH}/pluginsdk")
if(ARCH_BITS EQUAL 32)
    target_link_libraries(${PROJECT_NAME}
        "${DBG_LIB_PATH}/x32dbg.lib"
        "${DBG_LIB_PATH}/x32bridge.lib"
    )
else()
    target_link_libraries(${PROJECT_NAME}
        "${DBG_LIB_PATH}/x64dbg.lib"
        "${DBG_LIB_PATH}/x64bridge.lib"
    )
endif()

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
