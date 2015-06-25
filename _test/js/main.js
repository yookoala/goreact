
var CommentBox = React.createClass({
    render: (function () {
        return React.createElement("div", {
            className: "commentBox"
        }, "Hello, ", this.props.data.hello, "! I am a CommentBox.");
    })
});
