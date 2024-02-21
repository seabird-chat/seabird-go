package seabird

import (
	"fmt"

	"github.com/seabird-chat/seabird-go/pb"
)

type MessageBuilder struct {
	Blocks []*pb.Block
}

func NewMessageBuilder() *MessageBuilder {
	return &MessageBuilder{}
}

func (mb *MessageBuilder) AddText(text string) *MessageBuilder {
	mb.Blocks = append(mb.Blocks, &pb.Block{
		Inner: &pb.Block_Text{
			Text: &pb.TextBlock{
				Text: text,
			},
		},
	})
	return mb
}

func (mb *MessageBuilder) AddTextf(format string, args ...interface{}) *MessageBuilder {
	mb.Blocks = append(mb.Blocks, &pb.Block{
		Inner: &pb.Block_Text{
			Text: &pb.TextBlock{
				Text: fmt.Sprintf(format, args...),
			},
		},
	})
	return mb
}

func (mb *MessageBuilder) AddItalics(f func(*MessageBuilder)) *MessageBuilder {
	innerMb := NewMessageBuilder()
	f(innerMb)

	mb.Blocks = append(mb.Blocks, &pb.Block{
		Inner: &pb.Block_Italics{
			Italics: &pb.ItalicsBlock{
				Inner: innerMb.Blocks,
			},
		},
	})

	return mb
}

func (mb *MessageBuilder) AddBold(f func(*MessageBuilder)) *MessageBuilder {
	innerMb := NewMessageBuilder()
	f(innerMb)

	mb.Blocks = append(mb.Blocks, &pb.Block{
		Inner: &pb.Block_Bold{
			Bold: &pb.BoldBlock{
				Inner: innerMb.Blocks,
			},
		},
	})

	return mb
}

func (mb *MessageBuilder) AddUnderline(f func(*MessageBuilder)) *MessageBuilder {
	innerMb := NewMessageBuilder()
	f(innerMb)

	mb.Blocks = append(mb.Blocks, &pb.Block{
		Inner: &pb.Block_Underline{
			Underline: &pb.UnderlineBlock{
				Inner: innerMb.Blocks,
			},
		},
	})

	return mb
}

func (mb *MessageBuilder) AddStrikethrough(f func(*MessageBuilder)) *MessageBuilder {
	innerMb := NewMessageBuilder()
	f(innerMb)

	mb.Blocks = append(mb.Blocks, &pb.Block{
		Inner: &pb.Block_Strikethrough{
			Strikethrough: &pb.StrikethroughBlock{
				Inner: innerMb.Blocks,
			},
		},
	})

	return mb
}

func (mb *MessageBuilder) AddInlineCode(text string) *MessageBuilder {
	mb.Blocks = append(mb.Blocks, &pb.Block{
		Inner: &pb.Block_InlineCode{
			InlineCode: &pb.InlineCodeBlock{
				Text: text,
			},
		},
	})

	return mb
}

func (mb *MessageBuilder) AddFencedCode(info string, text string) *MessageBuilder {
	mb.Blocks = append(mb.Blocks, &pb.Block{
		Inner: &pb.Block_FencedCode{
			FencedCode: &pb.FencedCodeBlock{
				Info: info,
				Text: text,
			},
		},
	})

	return mb
}

func (mb *MessageBuilder) AddMention(user *pb.User) *MessageBuilder {
	mb.Blocks = append(mb.Blocks, &pb.Block{
		Inner: &pb.Block_Mention{
			Mention: &pb.MentionBlock{
				User: user,
			},
		},
	})

	return mb
}

func (mb *MessageBuilder) AddSpoiler(f func(*MessageBuilder)) *MessageBuilder {
	innerMb := NewMessageBuilder()
	f(innerMb)

	mb.Blocks = append(mb.Blocks, &pb.Block{
		Inner: &pb.Block_Spoiler{
			Spoiler: &pb.SpoilerBlock{
				Inner: innerMb.Blocks,
			},
		},
	})

	return mb
}

func (mb *MessageBuilder) AddAction(f func(*MessageBuilder)) *MessageBuilder {
	innerMb := NewMessageBuilder()
	f(innerMb)

	mb.Blocks = append(mb.Blocks, &pb.Block{
		Inner: &pb.Block_Action{
			Action: &pb.ActionBlock{
				Inner: innerMb.Blocks,
			},
		},
	})

	return mb
}

func (mb *MessageBuilder) AddList(f func(*MessageBuilder)) *MessageBuilder {
	innerMb := NewMessageBuilder()
	f(innerMb)

	mb.Blocks = append(mb.Blocks, &pb.Block{
		Inner: &pb.Block_List{
			List: &pb.ListBlock{
				Inner: innerMb.Blocks,
			},
		},
	})

	return mb
}

func (mb *MessageBuilder) AddLink(url string, f func(*MessageBuilder)) *MessageBuilder {
	innerMb := NewMessageBuilder()
	f(innerMb)

	mb.Blocks = append(mb.Blocks, &pb.Block{
		Inner: &pb.Block_Link{
			Link: &pb.LinkBlock{
				Url:   url,
				Inner: innerMb.Blocks,
			},
		},
	})

	return mb
}

func (mb *MessageBuilder) AddBlockquote(f func(*MessageBuilder)) *MessageBuilder {
	innerMb := NewMessageBuilder()
	f(innerMb)

	mb.Blocks = append(mb.Blocks, &pb.Block{
		Inner: &pb.Block_Blockquote{
			Blockquote: &pb.BlockquoteBlock{
				Inner: innerMb.Blocks,
			},
		},
	})

	return mb
}
