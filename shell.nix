{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  packages = [pkgs.go];
  inputsFrom = [pkgs.bat];

  shellHook = ''
    echo "welcome to my first development sshell!!"
  '';

  test = "DLKLDKSLKDL";
  envarr =  "SOMETESTING";

}

