

config.action_view.field_error_proc = Proc.new { |html_tag, instance|
  if instance.kind_of?(ActionView::Helpers::Tags::Label)
    html_tag.html_safe
  else
    %(#{html_tag}
      <my-form-error>#{instance.error_message.join(?,)}</my-form-error>
    ).html_safe
  end
}
