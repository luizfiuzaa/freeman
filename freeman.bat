@echo off
echo λ ~ FREEMAN WAS HERE! LET'S GO...

echo λ ~ FREEMAN DOING WHAT NEEDS TO BE DONE 

flutter clean
flutter pub cache clean
flutter pub cache repair

echo λ ~ FREEMAN DOING THE SPECIFIC CLEANUP!
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

echo λ ~ FREEMAN RELOADING DEPENDENCIES
flutter pub get

echo λ ~ FREEMAN IS OUT! HAVE A GREAT DAY!
