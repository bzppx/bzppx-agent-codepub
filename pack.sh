#!/bin/bash
VER=$1
if [ "$VER" = "" ]; then
    echo 'please input pack version!'
    exit 1
fi
RELEASE="release-${VER}"
rm -rf ${RELEASE}
mkdir ${RELEASE}

# windows amd64
echo 'Start pack windows amd64...'
GOOS=windows GOARCH=amd64 go build  
tar -czvf "${RELEASE}/bzppx-agent-codepub-windows-amd64.tar.gz" bzppx-agent-codepub.exe config.toml cert/ log/.gitignore LICENSE README.md
rm -rf bzppx-agent-codepub.exe

echo 'Start pack windows X386...'
GOOS=windows GOARCH=386 go build 
tar -czvf "${RELEASE}/bzppx-agent-codepub-windows-386.tar.gz" bzppx-agent-codepub.exe config.toml cert/ log/.gitignore LICENSE README.md
rm -rf bzppx-agent-codepub.exe

echo 'Start pack linux amd64'
GOOS=linux GOARCH=amd64 go build -ldflags "-s -w"
tar -czvf "${RELEASE}/bzppx-agent-codepub-linux-amd64.tar.gz" bzppx-agent-codepub config.toml cert/ log/.gitignore LICENSE README.md
rm -rf bzppx-agent-codepub

echo 'Start pack linux 386'
GOOS=linux GOARCH=386 go build -ldflags "-s -w"
tar -czvf "${RELEASE}/bzppx-agent-codepub-linux-386.tar.gz" bzppx-agent-codepub config.toml cert/ log/.gitignore LICENSE README.md
rm -rf bzppx-agent-codepub

echo 'Start pack mac amd64'
GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w"
tar -czvf "${RELEASE}/bzppx-agent-codepub-mac-amd64.tar.gz" bzppx-agent-codepub config.toml cert/ log/.gitignore LICENSE README.md
rm -rf bzppx-agent-codepub

echo 'END'
