package translation

import (
	"fmt"
	"time"

	"go.temporal.io/sdk/workflow"
)

func SayHelloGoodbye(ctx workflow.Context, input TranslationWorkflowInput) (TranslationWorkflowOutput, error) {
	// TODO define the Workflow logger here
	logger := workflow.GetLogger(ctx)

	// TODO Log, at the Info level, when the Workflow function is invoked
	//      and be sure to include the name passed as input
	logger.Info("SayHelloGoodbye invoked", "name", input.Name)

	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 45,
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	// TODO Log, at the Debug level, a message about the Activity to be executed,
	//      be sure to include the language code passed as input
	logger.Debug("translating hello message", "language code", input.LanguageCode)
	helloInput := TranslationActivityInput{
		Term:         "Hello",
		LanguageCode: input.LanguageCode,
	}
	var helloResult TranslationActivityOutput
	err := workflow.ExecuteActivity(ctx, TranslateTerm, helloInput).Get(ctx, &helloResult)
	if err != nil {
		return TranslationWorkflowOutput{}, err
	}
	helloMessage := fmt.Sprintf("%s, %s", helloResult.Translation, input.Name)

	// TODO: (Part C): log a message at the Debug level and then start a Timer for 10 seconds
	logger.Debug("starting timer")
	if err := workflow.NewTimer(ctx, 10*time.Second).Get(ctx, nil); err != nil {
		return TranslationWorkflowOutput{}, err
	}

	// TODO Log, at the Debug level, a message about the Activity to be executed,
	//      be sure to include the language code passed as input
	logger.Debug("translating goodbye message", "language code", input.LanguageCode)
	goodbyeInput := TranslationActivityInput{
		Term:         "Goodbye",
		LanguageCode: input.LanguageCode,
	}
	var goodbyeResult TranslationActivityOutput
	err = workflow.ExecuteActivity(ctx, TranslateTerm, goodbyeInput).Get(ctx, &goodbyeResult)
	if err != nil {
		return TranslationWorkflowOutput{}, err
	}
	goodbyeMessage := fmt.Sprintf("%s, %s", goodbyeResult.Translation, input.Name)

	output := TranslationWorkflowOutput{
		HelloMessage:   helloMessage,
		GoodbyeMessage: goodbyeMessage,
	}

	return output, nil
}
