## Install
* use go
    ```shell
    go get github.com/evgeny-klyopov/pack-mnemonic-app
    ```
* source
    * Mac OS (m1)
      ```shell
      curl -s https://api.github.com/repos/evgeny-klyopov/pack-mnemonic-app/releases/latest \
      | grep browser_download_url \
      | grep pm.macos-arm64.tar.gz \
      | cut -d '"' -f 4 \
      | wget -qi - 
      tar -xvf pm.macos-arm64.tar.gz && mv pm /usr/bin/pm 
      ```
    * Mac OS
      ```shell
      curl -s https://api.github.com/repos/evgeny-klyopov/pack-mnemonic-app/releases/latest \
      | grep browser_download_url \
      | grep pm.macos-amd64.tar.gz \
      | cut -d '"' -f 4 \
      | wget -qi - 
      tar -xvf pm.macos-amd64.tar.gz && mv pm /usr/bin/pm 
      ```
    * Linux
      ```shell
      curl -s https://api.github.com/repos/evgeny-klyopov/pack-mnemonic-app/releases/latest \
      | grep browser_download_url \
      | grep pm.linux-amd64.tar.gz \
      | cut -d '"' -f 4 \
      | wget -qi - 
      tar -xvf pm.linux-amd64.tar.gz && mv pm /usr/bin/pm 
      ```
    * Windows
      ```shell
      curl -s https://api.github.com/repos/evgeny-klyopov/pack-mnemonic-app/releases/latest \
      | grep browser_download_url \
      | grep pm.windows-amd64.tar.gz \
      | cut -d '"' -f 4 \
      | wget -qi -
      tar -xvf pm.windows-amd64.tar.gz && mv pm.exe /usr/bin/pm
      ```