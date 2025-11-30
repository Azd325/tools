{
  description = "My little tools collection";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs =
    {
      self,
      nixpkgs,
      flake-utils,
    }:
    flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = nixpkgs.legacyPackages.${system};

        toolNames = [
          "browser-cookies"
          "browser-eval"
          "browser-nav"
          "browser-screenshot"
          "browser-start"
          "browser-stop"
        ];

        toolPackages = builtins.listToAttrs (
          map (name: {
            inherit name;
            value = pkgs.callPackage ./${name} { };
          }) toolNames
        );

      in
      {
        packages = toolPackages // {
          default = toolPackages.browser-start;
        };

        devShells.default = pkgs.mkShell {
          buildInputs = [
            pkgs.go
            pkgs.gopls
          ];

          shellHook = ''
            echo "Browser Tools development shell"
            echo "Go: $(go version)"
          '';
        };

        formatter = pkgs.nixfmt-rfc-style;

      }
    )
    // {
      overlays.default =
        final: prev:
        let
          system = prev.stdenv.hostPlatform.system;
        in
        self.packages.${system} or { };
    };
}
