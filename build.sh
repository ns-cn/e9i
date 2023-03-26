#!/bin/zsh

build=20230326

rm -rf e9i.app && fyne package -os darwin --appVersion "1.0.0" --appBuild ${build} --release true -appID com.tangyujun.e9i -name "E9I" .
fyne package -os linux --appVersion "1.0.0" --appBuild ${build} --release true -appID com.tangyujun.e9i -name "E9I" .
fyne package -os windows --appVersion "1.0.0" --appBuild ${build} --release true -appID com.tangyujun.e9i -name "E9I" .
