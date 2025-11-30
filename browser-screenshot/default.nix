{ buildGoModule, lib }:

buildGoModule {
  pname = "browser-screenshot";
  version = "1.0.0";
  src = ./.;
  vendorHash = "sha256-xPoa/axatOR5v1oPwDZVo6r7SnFqehLI2zSm6EyIfkk=";

  meta = {
    description = "Capture browser screenshots";
    mainProgram = "browser-screenshot";
    license = lib.licenses.mit;
  };
}
