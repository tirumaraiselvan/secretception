# secretception

A K8s custom controller that encrypts a secret

This is a very stupid controller that takes a secret and md5 hashes it and stores it back.

This was used for DEMO purposes in Koffee With Kubernetes (https://twitter.com/KoffeeWithK8S) meetup on 8 March 2018

## Usage

go build

./secretception --kubeconfig <your-kubeconfig-loc>
