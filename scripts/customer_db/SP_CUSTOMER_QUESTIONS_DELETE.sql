SET ANSI_NULLS ON;
SET QUOTED_IDENTIFIER ON;
CREATE OR ALTER PROCEDURE [dbo].[SP_CUSTOMER_QUESTIONS_DELETE] (
	@customerId VARCHAR (64), @questionId int, @userId VARCHAR (64)
) AS
BEGIN
	SET ANSI_NULLS ON;
	SET ANSI_PADDING ON;
	SET QUOTED_IDENTIFIER ON;
	SET NOCOUNT ON;

	UPDATE [dbo].[CUSTOMER_QUESTIONS]
		SET [isDeleted] = 1, [updatedBy] = @userId, [updatedOn] = GETDATE()
		WHERE [customerId] = @customerId AND [questionId] = @questionId;
END;
-- exec [dbo].[SP_CUSTOMER_QUESTIONS_DELETE] 'SUPERUSER1', 2, 'TEST_USER4'