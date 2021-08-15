---
title: kugo
---

# Kubernetes client-go sample

* **Kugo**: Một chiếc `cli tool` nho nhỏ dành cho kubernetes sử  dụng thư viện
[client-go](https://github.com/kubernetes/client-go).
* **Kugo** giúp bạn `update image` của ứng dụng đang triển khai.

## Pre-Install

* Cài đặt `Golang`
* Cụm kubernetes để  triển khai

## Install

```bash
git clone https://github.com/ductnn/kugo.git
cd kugo
go build main.go
```

## Usage

* Sau khi `build` chương trình sinh ra file `main`, và để  sử  dụng cần truyền
vào các `arguments` sau:
    * `-deployment name:<string>`: Tên của `deployment` đang chạy.
    * `-app name:<string>`: Tên của ứng dụng đang chạy (default: `app`).
    * `-image name:<string>`: Tên image mới cần deploy.
    * `-kubeconfig string`: Đường dẫn file `kubeconfig` (default: `/home/(userid)/.kube/config`)

```bash
➜  kugo git:(master) ✗ ./main -deployment test-app -image nginx:1.13 -app test-app
```

## Example

* Mình đang chạy `deployment` có tên là `test-app` và sử  dụng `image: ductn4/green-rain:v1`,
và bây giờ mình thay đổi phiên bản `image` đang dùng thành `image: ductn4/green-rain:v2`

```bash
➜  kugo git:(master) ✗ ./main -deployment test-app -image ductn4/green-rain:v2 -app test-app
/home/ductn/.kube/config
Found deployment
name -> test-app
Old image -> ductn4/green-rain:v1
New image -> ductn4/green-rain:v2
```

[me](https://ductn.info/about)
