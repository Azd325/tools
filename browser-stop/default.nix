{
  buildGoModule,
  lib,
  makeWrapper,
  procps,
}:

buildGoModule {
  pname = "browser-stop";
  version = "1.0.0";
  src = ./.;
  vendorHash = null;

  nativeBuildInputs = [ makeWrapper ];

  postInstall = ''
    wrapProgram $out/bin/browser-stop \
      --prefix PATH : ${lib.makeBinPath [ procps ]}
  '';

  meta = {
    description = "Stop the CDP-controlled Chrome instance";
    mainProgram = "browser-stop";
    license = lib.licenses.mit;
  };
}
