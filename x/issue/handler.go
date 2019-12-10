package issue

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/handler"
	"github.com/konstellation/konstellation/x/issue/types"
)

func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				sdk.EventTypeMessage,
				sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
				//sdk.NewAttribute(sdk.AttributeKeySender, msg.GetSigners()[0].String()),
			),
		)

		switch msg := msg.(type) {
		case types.MsgIssue:
			return handler.HandleMsgIssue(ctx, k, msg)
		case types.MsgTransfer:
			return handler.HandleMsgTransfer(ctx, k, msg)
		case types.MsgTransferFrom:
			return handler.HandleMsgTransferFrom(ctx, k, msg)
		case types.MsgApprove:
			return handler.HandleMsgApprove(ctx, k, msg)
		case types.MsgIncreaseAllowance:
			return handler.HandleMsgIncreaseAllowance(ctx, k, msg)
		case types.MsgDecreaseAllowance:
			return handler.HandleMsgDecreaseAllowance(ctx, k, msg)

		default:
			errMsg := fmt.Sprintf("unrecognized issue message type: %T", msg)
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}