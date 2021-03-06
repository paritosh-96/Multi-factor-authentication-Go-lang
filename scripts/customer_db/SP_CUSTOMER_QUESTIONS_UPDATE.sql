SET ANSI_NULLS ON;
SET QUOTED_IDENTIFIER ON;
CREATE OR ALTER PROCEDURE [dbo].[SP_CUSTOMER_QUESTIONS_UPDATE] (
	@customerId VARCHAR (64), @questionId int, @answer VARCHAR (512)
) AS
BEGIN
	SET ANSI_NULLS ON;
	SET ANSI_PADDING ON;
	SET QUOTED_IDENTIFIER ON;
	SET NOCOUNT ON;

	UPDATE [dbo].[CUSTOMER_QUESTIONS]
		SET [answer] = @answer, [updatedBy] = @customerId, [updatedOn] = GETDATE()
		WHERE [customerId] = @customerId AND [questionId] = @questionId;
END;
-- exec [dbo].[SP_CUSTOMER_QUESTIONS_UPDATE] 'SUPERUSER1', 3, 'SHADOW', 'TEST_USER3'