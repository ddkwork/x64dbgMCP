rmdir /s /q cmake-build-debug
cmake -G "Visual Studio 17 2022" -A x64 -S . -B cmake-build-debug
cmake --build cmake-build-debug --config Release -j 6
copy cmake-build-debug\Release\MCPx64dbg.dp64 .

rmdir /s /q cmake-build-debug
cmake -G "Visual Studio 17 2022" -A Win32 ^
  -T "host=x64" ^
  -S . -B cmake-build-debug
cmake --build cmake-build-debug --config Release -j 6
copy cmake-build-debug\Release\MCPx64dbg.dp32 .

rmdir /s /q cmake-build-debug