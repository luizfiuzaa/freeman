@echo off
echo FREEMAN WAS HERE! Cleaning Cache...

flutter clean
flutter pub clean cache
flutter pub cache repair
rmdir /S /Q build/
rmdir /S /Q .dart_tool/
rmdir /S /Q .gradle/
rmdir /S /Q .idea/
rmdir /S /Q .packages
rmdir /S /Q ios/Pods
rmdir /S /Q ios/.symlinks
rmdir /S /Q ios/Flutter/Flutter.framework
rmdir /S /Q ios/Flutter/Flutter.podspec
rmdir /S /Q ios/Flutter/App.framework
rmdir /S /Q android/.gradle
rmdir /S /Q android/.idea
rmdir /S /Q android/.gradle/caches/
rmdir /S /Q android/.gradle/daemon/
rmdir /S /Q android/.gradle/native/
rmdir /S /Q android/.gradle/7.0/
rmdir /S /Q android/build/
del /F /Q pubspec.lock

flutter pub get


echo Cache cleaned! Have a good day!

