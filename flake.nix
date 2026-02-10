{
  description = "go-distro - Go library for detecting OS and distributions";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
      in
      {
        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [
            # Go development
            go
            gopls
            golangci-lint
            go-tools
            delve

            # Docker for testing
            docker
            docker-compose

            # General utilities
            git
            gnumake
          ];

          shellHook = ''
            echo "go-distro development environment"
            echo "Go version: $(go version)"
            echo ""
            echo "Available commands:"
            echo "  go build ./...    - Build all packages"
            echo "  go test ./...     - Run all tests"
            echo "  make test-docker  - Run Docker-based tests"
          '';
        };
      }
    );
}
