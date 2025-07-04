cmake_minimum_required(VERSION 3.15)
project(MCPx64dbg)

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

set(pluginsdk_dir "${CMAKE_SOURCE_DIR}/pluginsdk")
set(project_root "${CMAKE_SOURCE_DIR}")

include_directories(
        "${project_root}"
        "${pluginsdk_dir}"
)

file(GLOB sources "*.cpp")

add_library(${PROJECT_NAME} SHARED ${sources})


if (CMAKE_SIZEOF_VOID_P EQUAL 8)
    file(GLOB_RECURSE LIBS
            ${pluginsdk_dir}/*_x64.lib
            ${pluginsdk_dir}/x64*.lib
    )
    set_target_properties(${PROJECT_NAME} PROPERTIES SUFFIX ".dp64")
else ()
    file(GLOB_RECURSE LIBS
            ${pluginsdk_dir}/*_x86.lib
            ${pluginsdk_dir}/x32*.lib
    )
    set_target_properties(${PROJECT_NAME} PROPERTIES SUFFIX ".dp32")
endif ()
message(STATUS "LIBS: ${LIBS}")

target_link_libraries(${PROJECT_NAME}
        ws2_32
        winhttp
        ${LIBS})

#copy MCPx64dbg.dp32  release\x32
#copy MCPx64dbg.dp64  release\x64



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
