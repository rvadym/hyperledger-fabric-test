package usecases

type BatchUpdateItemsInterface interface {
	ExecuteBatchUpdateItems(args []string) error
}

func NewBatchUpdateItems(updateItemUC UpdateItemInterface) BatchUpdateItemsInterface {
	return &batchUpdateItems{
		updateItemUC,
	}
}

type batchUpdateItems struct {
	updateItemUC UpdateItemInterface
}

func (uc *batchUpdateItems) ExecuteBatchUpdateItems(args []string) error {
	var err error

	argsLen := len(args)
	cycles := (argsLen + 1) / 2

	for i := 1; i < cycles; i++ {
		id := ""
		content := ""

		idPos := i*2 - 1
		contentPos := i * 2

		if idPos <= argsLen {
			id = args[idPos]
		}

		if contentPos <= argsLen {
			content = args[contentPos]
		}

		_, err = uc.updateItemUC.ExecuteUpdateItem(id, content)
		if err != nil {
			return err
		}
	}

	return nil
}
