{ buildGoModule, lib }:

buildGoModule {
  pname = "browser-cookies";
  version = "1.0.0";
  src = ./.;
  vendorHash = "sha256-xPoa/axatOR5v1oPwDZVo6r7SnFqehLI2zSm6EyIfkk=";

  meta = {
    description = "List browser cookies for current page";
    mainProgram = "browser-cookies";
    license = lib.licenses.mit;
  };
}
