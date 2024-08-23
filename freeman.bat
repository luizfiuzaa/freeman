@echo off

echo $ ~ FREEMAN WAS HERE! LET'S GO...
echo.
echo $ ~ FREEMAN DOING WHAT NEEDS TO BE DONE 
echo.
echo y | flutter clean
echo $ ~ 1/3
echo.
echo y | flutter pub cache repair
echo.
echo $ ~ 2/3
echo.
echo y | flutter pub cache clean
echo $ ~ 3/3
echo.
echo $ ~ FREEMAN DOING THE SPECIFIC CLEANUP!
rmdir /S /Q build/ 2>nul
rmdir /S /Q .dart_tool/ 2>nul
rmdir /S /Q .gradle/ 2>nul
rmdir /S /Q .idea/ 2>nul
rmdir /S /Q .packages 2>nul
rmdir /S /Q ios/Pods 2>nul
rmdir /S /Q ios/.symlinks 2>nul
rmdir /S /Q ios/Flutter/Flutter.framework 2>nul
rmdir /S /Q ios/Flutter/Flutter.podspec 2>nul
rmdir /S /Q ios/Flutter/App.framework 2>nul
rmdir /S /Q android/.gradle 2>nul
rmdir /S /Q android/.idea 2>nul 
rmdir /S /Q android/.gradle/caches/ 2>nul
rmdir /S /Q android/.gradle/daemon/ 2>nul
rmdir /S /Q android/.gradle/native/ 2>nul
rmdir /S /Q android/.gradle/7.0/ 2>nul
rmdir /S /Q android/build/ 2>nul
del /F /Q pubspec.lock 2>nul

echo $ ~ FREEMAN RELOADING DEPENDENCIES
echo.

echo y | flutter pub get

echo $ ~ FREEMAN IS OUT! HAVE A GREAT DAY!
