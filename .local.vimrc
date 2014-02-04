fu! GoTest()
	
	if @% !~ ".go"

		return

	elseif @% =~ "_test.go"

		let file = @%

	else
		let file = split(@%, ".go")[0] . "_test.go"

	endif

	" now run go tests
	let command = "go test " . file

	" now run this command
	call CleanShell(command)

endfunction

map <Leader>rr :call GoTest()<CR>

