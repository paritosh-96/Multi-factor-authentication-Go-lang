SET ANSI_NULLS ON;
SET QUOTED_IDENTIFIER ON;
GO
CREATE OR ALTER PROCEDURE [dbo].[SP_CUSTOMER_QUESTIONS_VALIDATE] (
	@customerId VARCHAR (64), @questionId int, @answer VARCHAR (512)
)
AS
BEGIN
	SET ANSI_NULLS ON;
	SET ANSI_PADDING ON;
	SET QUOTED_IDENTIFIER ON;
	SET NOCOUNT ON;

	declare @fetchedAnswer int;
	SELECT @fetchedAnswer = [answer]
	FROM [dbo].[CUSTOMER_QUESTIONS] WITH (NOLOCK)
	WHERE [isDeleted] = 0 AND [customerId] = @customerId AND [questionId] = @questionId;

	IF @fetchedAnswer = @answer
		BEGIN
			return 0;
			END;
	ELSE
		BEGIN
			return 1;
		END;

END;