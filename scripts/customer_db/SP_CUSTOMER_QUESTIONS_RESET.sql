SET ANSI_NULLS ON;
SET QUOTED_IDENTIFIER ON;
CREATE OR ALTER PROCEDURE [dbo].[SP_CUSTOMER_QUESTIONS_RESET] (
	@customerId VARCHAR (64), @userId VARCHAR (64)
) AS
BEGIN
	SET ANSI_NULLS ON;
	SET ANSI_PADDING ON;
	SET QUOTED_IDENTIFIER ON;
	SET NOCOUNT ON;

	UPDATE [dbo].[CUSTOMER_QUESTIONS]
		SET [isDeleted] = 1 , [updatedBy] = @userId, [updatedOn] = GETDATE()
		WHERE [customerId] = @customerId;
END;
-- select * from CUSTOMER_QUESTIONS
-- exec [dbo].[SP_CUSTOMER_QUESTIONS_RESET] 'SUPERUSER1', 'TEST_USER3'