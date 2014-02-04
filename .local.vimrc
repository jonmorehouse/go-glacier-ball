fu! GoTestPackage()

	let command = "go test ."

	call CleanShell(command)

endfunction

fu! GoTestFile()
	
	if @% !~ ".go"

		return

	elseif @% =~ "_test.go"

		let file = @%

	else
		let file = split(@%, ".go")[0] . "_test.go"

	endif

	" now run go tests
	let command = "go test -v " . file

	" now run this command
	call CleanShell(command)

endfunction

map <Leader>rr :call GoTestFile()<CR>
map <Leader>r :call GoTestPackage()<CR>


