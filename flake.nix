{
  outputs = {
    self,
    nixpkgs,
  }: {
    devShells.aarch64-darwin.default = let
      pkgs = import nixpkgs {system = "aarch64-darwin";};
      inherit (pkgs) mkShell go;
    in
      mkShell {
        name = "gatekeeper";
        packages = [go];
      };
  };
}
