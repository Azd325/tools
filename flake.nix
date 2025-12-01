{
  description = "My little tools collection";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
  };

  outputs =
    { self, nixpkgs }:
    let
      forAllSystems =
        function:
        nixpkgs.lib.genAttrs nixpkgs.lib.systems.flakeExposed (
          system: function nixpkgs.legacyPackages.${system}
        );

      toolNames = [
        "browser-cookies"
        "browser-eval"
        "browser-nav"
        "browser-screenshot"
        "browser-start"
        "browser-stop"
      ];

      toolPackages =
        pkgs:
        builtins.listToAttrs (
          map (name: {
            inherit name;
            value = pkgs.callPackage ./${name} { };
          }) toolNames
        );
    in
    {
      packages = forAllSystems (
        pkgs:
        (toolPackages pkgs)
        // {
          default = (toolPackages pkgs).browser-start;
        }
      );

      devShells = forAllSystems (pkgs: {
        default = pkgs.mkShell {
          buildInputs = [
            pkgs.go
            pkgs.gopls
          ];

          shellHook = ''
            echo "Browser Tools development shell"
            echo "Go: $(go version)"
          '';
        };
      });

      formatter = forAllSystems (pkgs: pkgs.nixfmt-rfc-style);

      overlays.default = final: prev: toolPackages prev;
    };
}
