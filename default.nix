{ stdenv, makeWrapper, lib, fetchFromGitHub, glibcLocales
, coreutils, curl, jshon, bc, gnused, perl, datamash, git }:

stdenv.mkDerivation rec {
  name = "setzer-mcd-${version}";
  version = "0.1.0";
  src = ./.;

  nativeBuildInputs = [makeWrapper];
  buildPhase = "true";
  makeFlags = ["prefix=$(out)"];
  postInstall = let path = lib.makeBinPath [
    coreutils curl jshon bc gnused perl datamash git
  ]; in ''
    wrapProgram "$out/bin/setzer" --set PATH "${path}" \
      ${if glibcLocales != null then
        "--set LOCALE_ARCHIVE \"${glibcLocales}\"/lib/locale/locale-archive"
        else ""}
  '';

  meta = with lib; {
    description = "Ethereum price feed tool";
    homepage = https://github.com/makerdao/setzer;
    license = licenses.gpl3;
    inherit version;
  };
}
