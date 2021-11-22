# shell.nix
# used in nixos to run make commands like nix-shell --run "make sqlc"
let
  pkgs = import <nixpkgs>  {};
  stdenv = pkgs.stdenv;
in
  stdenv.mkDerivation {
    name = "env";
    buildInputs = with pkgs; [ gnumake ];
}