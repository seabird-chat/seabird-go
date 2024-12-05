package seabird

import (
	"fmt"
	"time"

	"github.com/seabird-chat/seabird-go/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func NewTextBlock(text string) *pb.Block {
	return &pb.Block{
		Inner: &pb.Block_Text{
			Text: &pb.TextBlock{
				Text: text,
			},
		},
	}
}

func NewTextBlockf(format string, args ...any) *pb.Block {
	return NewTextBlock(fmt.Sprintf(format, args...))
}

func NewInlineCodeBlock(text string) *pb.Block {
	return &pb.Block{
		Inner: &pb.Block_InlineCode{
			InlineCode: &pb.InlineCodeBlock{
				Text: text,
			},
		},
	}
}

func NewItalicsBlock(inner ...*pb.Block) *pb.Block {
	return &pb.Block{
		Inner: &pb.Block_Italics{
			Italics: &pb.ItalicsBlock{
				Inner: inner,
			},
		},
	}
}
func NewBoldBlock(inner ...*pb.Block) *pb.Block {
	return &pb.Block{
		Inner: &pb.Block_Bold{
			Bold: &pb.BoldBlock{
				Inner: inner,
			},
		},
	}
}

func NewUnderlineBlock(inner ...*pb.Block) *pb.Block {
	return &pb.Block{
		Inner: &pb.Block_Underline{
			Underline: &pb.UnderlineBlock{
				Inner: inner,
			},
		},
	}
}

func NewStrikethroughBlock(inner ...*pb.Block) *pb.Block {
	return &pb.Block{
		Inner: &pb.Block_Strikethrough{
			Strikethrough: &pb.StrikethroughBlock{
				Inner: inner,
			},
		},
	}
}

func NewFencedCodeBlock(info string, text string) *pb.Block {
	return &pb.Block{
		Inner: &pb.Block_FencedCode{
			FencedCode: &pb.FencedCodeBlock{
				Info: info,
				Text: text,
			},
		},
	}
}

func NewMentionBlock() *pb.Block {
	return &pb.Block{
		Inner: &pb.Block_Mention{
			Mention: &pb.MentionBlock{
				User: nil,
			},
		},
	}
}

func NewSpoilerBlock(inner ...*pb.Block) *pb.Block {
	return &pb.Block{
		Inner: &pb.Block_Spoiler{
			Spoiler: &pb.SpoilerBlock{
				Inner: inner,
			},
		},
	}
}

func NewBlockquoteBlock(inner ...*pb.Block) *pb.Block {
	return &pb.Block{
		Inner: &pb.Block_Blockquote{
			Blockquote: &pb.BlockquoteBlock{
				Inner: inner,
			},
		},
	}
}

func NewContainerBlock(inner ...*pb.Block) *pb.Block {
	return &pb.Block{
		Inner: &pb.Block_Container{
			Container: &pb.ContainerBlock{
				Inner: inner,
			},
		},
	}
}

func NewLinkBlock(url string, inner ...*pb.Block) *pb.Block {
	return &pb.Block{
		Inner: &pb.Block_Link{
			Link: &pb.LinkBlock{
				Url:   url,
				Inner: inner,
			},
		},
	}
}

func NewTimestampBlock(target time.Time) *pb.Block {
	return &pb.Block{
		Inner: &pb.Block_Timestamp{
			Timestamp: &pb.TimestampBlock{
				Inner: timestamppb.New(target),
			},
		},
	}
}
