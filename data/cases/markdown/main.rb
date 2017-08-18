require 'redcarpet'

# Initializes a Markdown parser
markdown = Redcarpet::Markdown.new(Redcarpet::Render::HTML)

# You can use HEREDOC to be more flexiable with input
input = <<-HEREDOC
# Test H1
and some text here

## Test
Other text

- First
  - Second
HEREDOC

puts markdown.render(input)

# As result you will have following html
# -------------------------------------
# <h1>Test H1</h1>
# <p>and some text here</p>
# <h2>Test</h2>
# <p>Other text</p>
# <ul>
# 	<li>First
# 		<ul>
# 			<li>Second</li>
# 		</ul>
# 	</li>
# </ul>
# -------------------------------------
