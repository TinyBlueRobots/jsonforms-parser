package jsonforms

// Visitor defines the interface for visiting UI schema elements
type Visitor interface {
	VisitControl(*Control) error
	VisitVerticalLayout(*VerticalLayout) error
	VisitHorizontalLayout(*HorizontalLayout) error
	VisitGroup(*Group) error
	VisitCategorization(*Categorization) error
	VisitCategory(*Category) error
	VisitLabel(*Label) error
	VisitCustomElement(*CustomElement) error
}

// Walk traverses a UI schema element tree and calls the appropriate visitor methods
func Walk(element UISchemaElement, visitor Visitor) error {
	if element == nil {
		return nil
	}

	switch e := element.(type) {
	case *Control:
		return visitor.VisitControl(e)
	case *VerticalLayout:
		if err := visitor.VisitVerticalLayout(e); err != nil {
			return err
		}

		for _, child := range e.Elements {
			if err := Walk(child, visitor); err != nil {
				return err
			}
		}
	case *HorizontalLayout:
		if err := visitor.VisitHorizontalLayout(e); err != nil {
			return err
		}

		for _, child := range e.Elements {
			if err := Walk(child, visitor); err != nil {
				return err
			}
		}
	case *Group:
		if err := visitor.VisitGroup(e); err != nil {
			return err
		}

		for _, child := range e.Elements {
			if err := Walk(child, visitor); err != nil {
				return err
			}
		}
	case *Categorization:
		if err := visitor.VisitCategorization(e); err != nil {
			return err
		}

		for _, child := range e.Elements {
			if err := Walk(child, visitor); err != nil {
				return err
			}
		}
	case *Category:
		if err := visitor.VisitCategory(e); err != nil {
			return err
		}

		for _, child := range e.Elements {
			if err := Walk(child, visitor); err != nil {
				return err
			}
		}
	case *Label:
		return visitor.VisitLabel(e)
	case *CustomElement:
		if err := visitor.VisitCustomElement(e); err != nil {
			return err
		}

		for _, child := range e.Elements {
			if err := Walk(child, visitor); err != nil {
				return err
			}
		}
	}

	return nil
}

// BaseVisitor provides default no-op implementations for all visitor methods
// This allows users to embed BaseVisitor and only override methods they care about
type BaseVisitor struct{}

func (b *BaseVisitor) VisitControl(*Control) error                   { return nil }
func (b *BaseVisitor) VisitVerticalLayout(*VerticalLayout) error     { return nil }
func (b *BaseVisitor) VisitHorizontalLayout(*HorizontalLayout) error { return nil }
func (b *BaseVisitor) VisitGroup(*Group) error                       { return nil }
func (b *BaseVisitor) VisitCategorization(*Categorization) error     { return nil }
func (b *BaseVisitor) VisitCategory(*Category) error                 { return nil }
func (b *BaseVisitor) VisitLabel(*Label) error                       { return nil }
func (b *BaseVisitor) VisitCustomElement(*CustomElement) error       { return nil }
