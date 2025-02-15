{
  description = "Go project with Echo, templ, MySQL, GORM, Air, Docker Compose, Nix flake example";

  inputs.nixpkgs.url = "github:NixOS/nixpkgs";
  inputs.flake-utils.url = "github:numtide/flake-utils";

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
      in {
        devShell = pkgs.mkShell {
          buildInputs = [
            pkgs.go
            pkgs.docker
            pkgs.docker-compose
            pkgs.goPackages.air  # Air for live reload
            pkgs.mysql           # MySQL client for convenience
          ];
        };
      }
    );
}
