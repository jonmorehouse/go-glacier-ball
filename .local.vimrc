fu! GoCheckSuite()

	let command = "GOPATH=". g:basePath . " go test -gocheck.f ProcessorSuite"
	call CleanShell(command)

endfunction


fu! GoCheckFile()

	if @% !~ ".go"

		return

	elseif @% =~ "_test.go"

		let file = @%

	else
		let file = split(@%, ".go")[0] . "_test.go"

	endif

	" now run go tests
	let command = "GOPATH=". g:basePath . " go test -gocheck.v " . file . " test_bootstrap.go"

	" now run this command
	call CleanShell(command)
endfunction

map <Leader>rr :call GoCheckSuite()<CR>
"map <Leader>rr :call CleanShell("")
map <Leader>r :call GoCheckFile()<CR>
