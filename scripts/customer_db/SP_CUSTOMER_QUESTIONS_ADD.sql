SET ANSI_NULLS ON;
SET QUOTED_IDENTIFIER ON;
CREATE OR ALTER PROCEDURE [dbo].[SP_CUSTOMER_QUESTIONS_ADD] (
	@questionId int, @customerId VARCHAR (64), @answer VARCHAR (512), @userId VARCHAR(64)
)
AS
BEGIN
	SET ANSI_NULLS ON;
	SET ANSI_PADDING ON;
	SET QUOTED_IDENTIFIER ON;
	SET NOCOUNT ON;

	IF EXISTS(SELECT TOP 1 1 FROM [dbo].[QUESTION_BANK] WITH (NOLOCK) WHERE [id] = @questionId)
		BEGIN
			INSERT INTO [dbo].[CUSTOMER_QUESTIONS] ([questionId], [customerId], [answer], [createdBy], [updatedBy])
				VALUES (@questionId, @customerId, @answer, @userID, @userID);
		END;
	ELSE
		BEGIN
			THROW 51000, 'The Question id does not exist.', 1;
		end
END;

-- exec [dbo].[SP_CUSTOMER_QUESTIONS_ADD] 1, 'SUPERUSER1', 'Tommy', 'TEST_USER';
