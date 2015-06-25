
var CommentBox = React.createClass({
	render: function() {
		return (
			<div className="commentBox">
				Hello, {this.props.data}! I am a CommentBox.
			</div>
		);
	}
});
