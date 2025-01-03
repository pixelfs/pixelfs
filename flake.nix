{
  description = "PixelFS is a cross-device file system, Transfer files based on s3-protocol.";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = {
    self,
    nixpkgs,
    flake-utils,
    ...
  }: let
    pixelfsVersion =
      if (self ? shortRev)
      then self.shortRev
      else "dev";

    buildPixelModule = { pkgs, pname, subPackage }: pkgs.buildGo123Module rec {
      inherit pname;
      CGO_ENABLED=0;
      version = pixelfsVersion;
      src = pkgs.lib.cleanSource self;

      # Only run unit tests when testing a build
      checkFlags = ["-short"];

      # When updating go.mod or go.sum, a new sha will need to be calculated,
      # update this if you have a mismatch after doing a change to thos files.
      vendorHash = "sha256-tMmevXD5uRo7hU64qftzQisNJCfedu0ZEStsFwhFEGE=";

      subPackages = [subPackage];

      ldflags = ["-s" "-w" "-X github.com/pixelfs/pixelfs/${subPackage}/cli.Version=v${version}"];
    };
  in
    {
      overlay = _: prev: let
        pkgs = nixpkgs.legacyPackages.${prev.system};
      in rec {
        pixelfs = buildPixelModule {
          pkgs = pkgs;
          pname = "pixelfs";
          subPackage = "cmd/pixelfs";
        };
      };
    }

    // flake-utils.lib.eachDefaultSystem
    (system: let
      pkgs = import nixpkgs {
        overlays = [self.overlay];
        inherit system;
      };
      buildDeps = with pkgs; [git go_1_23 gnumake];
      devDeps = with pkgs;
        buildDeps
        ++ [
          golangci-lint
          golines
          nodePackages.prettier
          goreleaser
          nfpm
          gotestsum
          gotests
          ksh
          ko
          yq-go
          ripgrep
          air
          gofumpt

          # Protobuf dependencies
          protobuf
          protoc-gen-go
          protoc-gen-go-grpc
          protoc-gen-connect-go
          buf
          clang-tools # clang-format
        ];

    in rec {
      # `nix develop`
      devShell = pkgs.mkShell {
        buildInputs =
          devDeps
          ++ [
            (pkgs.writeShellScriptBin
              "go-mod-update-all"
              ''
                cat go.mod | ${pkgs.silver-searcher}/bin/ag "\t" | ${pkgs.silver-searcher}/bin/ag -v indirect | ${pkgs.gawk}/bin/awk '{print $1}' | ${pkgs.findutils}/bin/xargs go get -u
                go mod tidy
              '')
          ];

        shellHook = ''
          export PATH="$PWD/result/bin:$PATH"
        '';
      };

      # `nix build`
      packages = with pkgs; {
        inherit pixelfs;
      };
      defaultPackage = pkgs.pixelfs;

      checks = {
        format =
          pkgs.runCommand "check-format"
          {
            buildInputs = with pkgs; [
              gnumake
              nixpkgs-fmt
              golangci-lint
              nodePackages.prettier
              golines
              clang-tools
            ];
          } ''
            ${pkgs.nixpkgs-fmt}/bin/nixpkgs-fmt ${./.}
            ${pkgs.golangci-lint}/bin/golangci-lint run --fix --timeout 10m
            ${pkgs.nodePackages.prettier}/bin/prettier --write '**/**.{ts,js,md,yaml,yml,sass,css,scss,html}'
            ${pkgs.golines}/bin/golines --max-len=88 --base-formatter=gofumpt -w ${./.}
            ${pkgs.clang-tools}/bin/clang-format -style="{BasedOnStyle: Google, IndentWidth: 4, AlignConsecutiveDeclarations: true, AlignConsecutiveAssignments: true, ColumnLimit: 0}" -i ${./.}
          '';
      };
    });
}
