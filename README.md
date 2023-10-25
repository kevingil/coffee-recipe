# Coffee Recipe with ChatGPT
Simple OpenAI completion app to generate a coffee bean recipe.

Send a prompt using a form submit and stream the response with SSE.

The prompt generates a simple recipe approximating grind amount and yield.

### Build 

You must setup an <code>.env</code> file first.

To get a key, go here: [openai.com/blog/openai-api](https://openai.com/blog/openai-api)
<pre>
PORT=8080
OPENAI_API_KEY=your_openai_key
</pre>

Download modules and run
<pre>
go mod tidy
go run main.go
</pre>


### Dependencies 

github.com/sashabaranov/go-openai


### Screenshots


![request](https://s6.imgcdn.dev/9I2zy.png)


![result](https://s6.imgcdn.dev/9IAf8.png)
