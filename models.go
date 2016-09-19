package psapi

type Category struct {
	ID int `xml:"id"`
	IDParent string `xml:"id_parent"`
	LevelDepth string `xml:"level_depth"`
	Active string `xml:"active"`
	IDShopDefault string `xml:"id_shop_default"`
	IsRootCategory string `xml:"is_root_category"`
	Position string `xml:"position"`
	DateAdd string `xml:"date_add"`
	DateUpd string `xml:"date_upd"`
	Name string `xml:"name"`
	LinkRewrite string `xml:"link_rewrite"`
	Description string `xml:"description"`
	MetaTitle string `xml:"meta_title"`
	MetaDescription string `xml:"meta_description"`
	MetaKeywords string `xml:"meta_keywords"`
	Associations struct {
		Categories []struct {
			ID string `xml:"id"`
		} `xml:"categories"`
	} `xml:"associations"`
}

type Product struct {
	ID string  `xml:"id,attr"`
	IDManufacturer string `xml:"id_manufacturer"`
	IDSupplier string `xml:"id_supplier"`
	IDCategoryDefault string `xml:"id_category_default"`
	New interface{} `xml:"new"`
	CacheDefaultAttribute string `xml:"cache_default_attribute"`
	IDDefaultImage string `xml:"id_default_image"`
	IDDefaultCombination string `xml:"id_default_combination"`
	IDTaxRulesGroup string `xml:"id_tax_rules_group"`
	PositionInCategory string `xml:"position_in_category"`
	ManufacturerName string `xml:"manufacturer_name"`
	Quantity string `xml:"quantity"`
	Type string `xml:"type"`
	IDShopDefault string `xml:"id_shop_default"`
	Reference string `xml:"reference"`
	Sku string `xml:"sku"`
	SupplierReference string `xml:"supplier_reference"`
	Location string `xml:"location"`
	Width string `xml:"width"`
	Height string `xml:"height"`
	Depth string `xml:"depth"`
	Weight string `xml:"weight"`
	QuantityDiscount string `xml:"quantity_discount"`
	Ean13 string `xml:"ean13"`
	Upc string `xml:"upc"`
	CacheIsPack string `xml:"cache_is_pack"`
	CacheHasAttachments string `xml:"cache_has_attachments"`
	IsVirtual string `xml:"is_virtual"`
	OnSale string `xml:"on_sale"`
	OnlineOnly string `xml:"online_only"`
	Ecotax string `xml:"ecotax"`
	MinimalQuantity string `xml:"minimal_quantity"`
	Price string `xml:"price"`
	WholesalePrice string `xml:"wholesale_price"`
	Unity string `xml:"unity"`
	UnitPriceRatio string `xml:"unit_price_ratio"`
	AdditionalShippingCost string `xml:"additional_shipping_cost"`
	Customizable string `xml:"customizable"`
	TextFields string `xml:"text_fields"`
	UploadableFiles string `xml:"uploadable_files"`
	Active string `xml:"active"`
	RedirectType string `xml:"redirect_type"`
	IDProductRedirected string `xml:"id_product_redirected"`
	AvailableForOrder string `xml:"available_for_order"`
	AvailableDate string `xml:"available_date"`
	Condition string `xml:"condition"`
	ShowPrice string `xml:"show_price"`
	Indexed string `xml:"indexed"`
	Visibility string `xml:"visibility"`
	AdvancedStockManagement string `xml:"advanced_stock_management"`
	DateAdd string `xml:"date_add"`
	DateUpd string `xml:"date_upd"`
	PackStockType string `xml:"pack_stock_type"`
	MetaDescription string `xml:"meta_description"`
	MetaKeywords string `xml:"meta_keywords"`
	MetaTitle string `xml:"meta_title"`
	LinkRewrite string `xml:"link_rewrite"`
	Name string `xml:"name"`
	Description string `xml:"description"`
	DescriptionShort string `xml:"description_short"`
	AvailableNow string `xml:"available_now"`
	AvailableLater string `xml:"available_later"`
	Associations struct {
		Categories []struct {
			ID string `xml:"id"`
		} `xml:"categories"`
		Images []struct {
			ID string `xml:"id"`
		} `xml:"images"`
		Combinations []struct {
			ID string `xml:"id"`
		} `xml:"combinations"`
		ProductOptionValues []struct {
			ID string `xml:"id"`
		} `xml:"product_option_values"`
		StockAvailables []struct {
			ID string `xml:"id"`
			IDProductAttribute string `xml:"id_product_attribute"`
		} `xml:"stock_availables"`
	} `xml:"associations"`
}