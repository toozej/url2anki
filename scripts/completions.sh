#!/bin/sh
set -e
rm -rf completions
mkdir completions
for sh in bash zsh fish; do
	go run ./cmd/url2anki/ completion "$sh" >"completions/url2anki.$sh"
done
