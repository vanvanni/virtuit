# VirtuIT

A minimal HTTP-based management layer for [Firecracker MicroVMs](https://firecracker-microvm.github.io/), designed as a development playground and experimentation platform. Ideal for testing and prototyping isolated environments with ease.

## Why did I create this?

I built VirtuIT as a foundation and testing ground for concepts behind a more complex, proprietary system I'm developing. A durable, workflow-based service that heavily relies on Firecracker for lightweight virtualization. Rather than keeping everything closed, I wanted to open-source a subset of that work to share how I manage and interact with Firecracker.

This project serves as a proof-of-concept: it represents roughly 25% of the production-grade system and acts as an experimental layer where I validate architectural and interface-level decisions.

## What is this repository?

This repository contains a simple HTTP-based hypervisor that can:

- Launch and manage Firecracker MicroVMs
- Interact with basic VM lifecycle operations via HTTP APIs
- Serve as a sandbox for experimenting with virtualized development environments

It is not intended for production use, but instead provides a transparent look at how to maybe self-manage Firecracker instances.

## Disclaimer

VirtuIT is a personal side project and is not meant to be used in production environments. It is experimental in nature and serves primarily as a playground and prototype for testing ideas related to Firecracker integration. Use it at your own risk.
