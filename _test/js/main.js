
var CommentBox = React.createClass({
    render: (function () {
        return React.createElement("div", {
            className: "commentBox"
        }, "Hello, ", this.props.data, "! I am a CommentBox.");
    })
});
