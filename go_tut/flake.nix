
{
  description = "Golang dev environment";

  inputs = {
    nixpkgs.url      = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url  = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils, ... }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs {
          inherit system;
        };
      in
      {
        devShells.default = with pkgs; mkShell {
          buildInputs = [
            bashInteractive
            bash-completion
            go
            gopls
            gofumpt
            golines
            gotools
            golangci-lint
            delve
          ];
          

          # Use zsh so it works with my neovim shell prompt
          shell = pkgs.bashInteractive;
          shellHook = ''
            echo "hello golang  "
          '';
        };
      }
    );
}
