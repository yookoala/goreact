
var CommentBox = React.createClass({
	render: function() {
		return (
			<div className="commentBox">
				Hello, {this.props.data.hello}! I am a CommentBox.
			</div>
		);
	}
});
