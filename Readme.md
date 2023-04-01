JSON-GO-Converter
=====

A simple program to automate converting any raw json to Golang structure.

Things to note:

- For Array in JSON every value has to be define in a new line.
- Converter decodes JSON line by line Key and value should be define in single line
- Params can be used to give name for main Go-package[<em>PackageName</em>] and Go-Structure name[<em>StructName</em>]