# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class LazyPublish < Formula
  desc ""
  homepage "https://github.com/yigitozgumus/lazy-publish"
  version "1.0.3"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/yigitozgumus/lazy-publish/releases/download/v1.0.3/lazy-publish_1.0.3_macOS_arm64.tar.gz"
      sha256 "c638b1a38a5da1dc4c4f188ed7ab370352bf130e17af04093430517d1172fc12"

      def install
        bin.install "lazy-publish"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/yigitozgumus/lazy-publish/releases/download/v1.0.3/lazy-publish_1.0.3_macOS_x86_64.tar.gz"
      sha256 "eb5465dd65dae04c542edacfce5deb8c354dd1e0ffff72b130cff2270a293b50"

      def install
        bin.install "lazy-publish"
      end
    end
  end

  on_linux do
    if Hardware::CPU.intel?
      url "https://github.com/yigitozgumus/lazy-publish/releases/download/v1.0.3/lazy-publish_1.0.3_Linux_x86_64.tar.gz"
      sha256 "17f713720f170990f9b9720443b27500fa36d4c5983edf55e9b5d2ce62946b51"

      def install
        bin.install "lazy-publish"
      end
    end
  end

  depends_on "git"
  depends_on "go"
end
