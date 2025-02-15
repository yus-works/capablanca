{
  description = "Go project with Echo, templ, MySQL, GORM, Air, Docker Compose, Nix flake example";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs";
    flake-utils.url = "github:numtide/flake-utils";
    templ.url = "github:a-h/templ";
  };

  outputs = { self, nixpkgs, flake-utils, templ }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs {
          inherit system;
          overlays = [ templ.overlays.default ];
        };
      in {
        devShell = pkgs.mkShell {
          buildInputs = [
            pkgs.go
            pkgs.docker
            pkgs.docker-compose
            pkgs.air             # Air for live reload
            pkgs.mariadb         # MySQL client for convenience
            pkgs.templ           # templ binary
          ];

          shellHook = ''
            export DATABASE_URL="mysql://capablanca:secret@tcp(127.0.0.1:3306)/capablanca"
            alias dc="docker-compose up -d"
            alias ga="git add"
            alias gs="git status"
            alias gc="git commit"
            alias gp="git push"
            echo ""
            echo "Added git aliases"
            echo "Development environment activated!"
            echo "Run 'dc' to start the db"
          '';
        };
      }
    );
}

