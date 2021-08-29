with (import <nixpkgs> { });
let
  packages = [
    go
    go-ethereum
  ];

in pkgs.mkShell {
  name = "ketherstatic";
  buildInputs = packages;
}
