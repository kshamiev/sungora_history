define(
	"main",
	[
		"MessageList"
	],
	function(MessageList) {
		var ws = new WebSocket("ws://localhost:8080/api/v1/websocket/gorilla");
		var list = new MessageList(ws);
		ko.applyBindings(list);
	}
);
