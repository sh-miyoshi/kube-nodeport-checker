# kube-nodeport-checker

## Overview

Kubernetes Cluster上で使用されているNodePortを表示します。  
このツールによってKubernetesのServiceの`type: NodePort`で使用しているポート番号がわかります。  
出力例はUsageの項目を参照してください。  

## How to Install

1. go言語がインストール済みの場合
    - godep get
    - make go\_build
    - make install
    - kube-nodeport-checker --help

2. Dockerがインストール済みの場合
    - make docker\_build
    - docker run --rm -it -v ~/.kube/config:/root/.kube/config kube-nodeport-checker:latest kube-nodeport-checker --help

## Usage

基本的にkube-nodeport-checkerコマンドを実行すると使用されているNodePortの一覧が出力されます。

```text
[root@server ~]# kube-nodeport-checker

30000: http (in nginx)
30010: no name (in wekan)
```

-kubeconfigオプションを指定した場合、kubernetes clusterにアクセスするためのconfigファイルの場所を指定できます。(指定しない場合は`~/.kube.config`ファイルが参照されます。)

## Author

Shunsuke Miyoshi
