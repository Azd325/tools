{
  buildGoModule,
  lib,
  makeWrapper,
  rsync,
}:

buildGoModule {
  pname = "browser-start";
  version = "1.0.0";
  src = ./.;
  vendorHash = null;

  nativeBuildInputs = [ makeWrapper ];

  postInstall = ''
    wrapProgram $out/bin/browser-start \
      --prefix PATH : ${lib.makeBinPath [ rsync ]}
  '';

  meta = {
    description = "Launch Chrome with CDP enabled";
    mainProgram = "browser-start";
    license = lib.licenses.mit;
  };
}
