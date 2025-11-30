{ buildGoModule, lib }:

buildGoModule {
  pname = "browser-nav";
  version = "1.0.0";
  src = ./.;
  vendorHash = "sha256-xPoa/axatOR5v1oPwDZVo6r7SnFqehLI2zSm6EyIfkk=";

  meta = {
    description = "Navigate browser to URL";
    mainProgram = "browser-nav";
    license = lib.licenses.mit;
  };
}
