{ buildGoModule, lib }:

buildGoModule {
  pname = "browser-eval";
  version = "1.0.0";
  src = ./.;
  vendorHash = "sha256-xPoa/axatOR5v1oPwDZVo6r7SnFqehLI2zSm6EyIfkk=";

  meta = {
    description = "Evaluate JavaScript in browser";
    mainProgram = "browser-eval";
    license = lib.licenses.mit;
  };
}
